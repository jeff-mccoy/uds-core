// Copyright 2026 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
import yaml from 'js-yaml';

export function validateRender(source,identityIncluded) {
  // Original gateway templates repeat the app/istio labels. Match the Helm
  // YAML-to-JSON conversion while preserving the original signed chart bytes.
  const objects=yaml.loadAll(source,null,{json:true}).filter(o=>o?.kind);
  const runtime=objects.filter(o=>['Deployment','DaemonSet','StatefulSet','Pod'].includes(o.kind)).map(o=>({kind:o.kind,name:o.metadata.name,namespace:o.metadata.namespace,images:(o.kind==='Pod'?o.spec:o.spec.template.spec).containers.map(c=>c.image)}));
  if(runtime.some(o=>o.images.some(i=>i.includes('###'))))throw new Error('Runtime image still contains an unresolved package template');
  if(runtime.some(o=>o.images.some(i=>/pepr|keycloak:|identity-config:/.test(i))))throw new Error('Excluded legacy runtime remains in native manifests');
  const controllers=objects.filter(o=>o.kind==='Deployment'&&o.metadata.name==='uds-controller');
  if(controllers.length!==4)throw new Error('Expected exactly four native controller installation phases');
  const configs=objects.filter(o=>['MutatingWebhookConfiguration','ValidatingWebhookConfiguration'].includes(o.kind)&&o.metadata.name.startsWith('uds-controller-'));
  if(configs.length!==15)throw new Error('Complete callbacks must occur only in registration, active and narrowed phases');
  const callbacks=configs.reduce((total,config)=>total+config.webhooks.length,0);
  if(callbacks!==27)throw new Error('Expected nine complete callback entries in each of three registered phases');
  for(const config of configs)for(const hook of config.webhooks)if(hook.failurePolicy!=='Fail'||!hook.clientConfig.caBundle)throw new Error('Callback registration lost fail-closed shared trust');
  const bindings=objects.filter(o=>o.kind.endsWith('AdmissionPolicyBinding')&&o.metadata.name.startsWith('uds-native-'));
  if(bindings.length!==28)throw new Error('Wrong native activation binding counts');
  let identityPasswordOnly=false;
  const secret=objects.find(o=>o.kind==='Secret'&&o.metadata.name==='uds-native-identity-config');
  if(identityIncluded){
    if(!secret)throw new Error('Native development identity configuration missing');
    const bridge=JSON.parse(secret.stringData['bridge.json']);
    if(bridge.passwordOnly!==true)throw new Error('Identity password-only contract omitted');
    const stateful=objects.find(o=>o.kind==='StatefulSet'&&o.metadata.namespace==='keycloak');
    if(stateful.spec.replicas!==1||stateful.spec.template.spec.serviceAccountName!=='uds-native-identity')throw new Error('Unqualified identity replica or authority topology');
    const dex=stateful.spec.template.spec.containers.filter(c=>c.name.startsWith('dex-'));
    const namespaces=dex.map(c=>c.env.find(e=>e.name==='KUBERNETES_POD_NAMESPACE')?.value);
    if(dex.length!==2||namespaces.some(n=>!n)||new Set(namespaces).size!==2)throw new Error('Public and admin Dex storage authority was combined');
    identityPasswordOnly=true;
  }else if(secret||runtime.some(o=>['keycloak','authservice','istio-admin-gateway','istio-tenant-gateway'].includes(o.namespace)))throw new Error('Identity or public gateways leaked into fenced native base');
  return {objects:objects.length,runtime,controllerPhases:controllers.length,completeCallbackRegistrations:configs.length,completeCallbackEntries:callbacks,nativeBindingsByName:Object.fromEntries([...new Set(bindings.map(o=>o.metadata.name))].map(name=>[name,bindings.filter(o=>o.metadata.name===name).length])),identityPasswordOnly,yamlDuplicateHandling:'Inherited original gateway labels use Helm JSON last-value behavior; signed source templates preserved'};
}
