# compute-node-operator
Compute-Node Operator


## Pre Req:
- OCP 4 installed
- operator-sdk 1.x

## Clone it

    git clone https://github.com/openstack-k8s-operators/compute-node-operator.git
    cd compute-node-operator

## Create the operator

This is optional, a prebuild operator from quay.io/ltomasbo/compute-operator could be used, e.g. quay.io/ltomasbo/compute-node-operator:v0.0.1 .

Build the image, using your custom registry you have write access to

    WATCH_NAMESPACE="openstack,openshift-machine-api" make docker-build IMG=<image e.g quay.io/ltomasbo/compute-node-operator:v0.0.X>

Push image to your registry:

    podman push --authfile ~/ltomasbo-auth.json quay.io/ltomasbo/compute-node-operator:v0.0.X

## Pre Req to run the operator

Requirement are the ConfigMap and Secret for the OpenStack client to be able to connect to the OpenStack API.
This is required for the pre/post scripts during scale down operation. In the end this should be created by a
control-plane operator.

Create `openstackclient` ConfigMap file (e.g. `openstackclient-cm.yaml`)

    apiVersion: v1
    data:
      OS_CLOUD: default
      clouds.yaml: |
        clouds:
          default:
            auth:
              auth_url: http://keystone-openstack.apps.ostest.test.metalkube.org
              project_name: admin
              username: admin
              user_domain_name: Default
              project_domain_name: Default
            region_name: regionOne
    kind: ConfigMap
    metadata:
      name: openstack-config
      namespace: openstack

Create `openstackclient-admin` Secret which for a user with admin privileges (e.g. `openstackclient-admin-secret.yaml`)

    apiVersion: v1
    data:
      secure.yaml: Y2xvdWRzOgogIGRlZmF1bHQ6CiAgICBhdXRoOgogICAgICBwYXNzd29yZDogZm9vYmFyMTIzCg==
    kind: Secret
    metadata:
      name: openstack-config-secret
      namespace: openstack

Create both, ConfigMap and Secret using:

    oc apply -f openstackclient-cm.yaml -f openstackclient-admin-secret.yaml

## Run the operator local for dev

    WATCH_NAMESPACE="openstack,openshift-machine-api" make run

## Install the operator to an OCP env

The operator is intended to be deployed via OLM [Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager)

If necessary check logs with

    POD=`oc get pods -l name=compute-node-operator --field-selector=status.phase=Running -o name | head -1 -`; echo $POD
    oc logs $POD -f

## Create a compute node CR

Create custom resource for a compute node which specifies the needed information (e.g.: `config/samples/compute-node_v1alpha1_computenodeopenstack.yaml`):

    apiVersion: compute-node.openstack.org/v1alpha1
    kind: ComputeNodeOpenStack
    metadata:
      name: example-computenodeopenstack
      namespace: openstack
    spec:
      # Add fields here
      roleName: worker-osp
      clusterName: ostest
      baseWorkerMachineSetName: ostest-worker-0
      workers: 0
      dedicated: false
      selinuxDisabled: true
      compute:
        novaComputeCPUDedicatedSet: "4-7"
        novaComputeCPUSharedSet: "0-3"
        sshdPort: 2022
      network:
        nic: "enp2s0"
        bridgeMappings: ""
        mechanismDrivers: ""
        servicePlugins: ""
        sriov:
          devName: "test"
    # needed if openshift-sdn is not running
    #  infraDaemonSets:
    #  - name: multus
    #    namespace: openshift-multus
    #  - name: node-exporter
    #    namespace: openshift-monitoring
    #  - name: machine-config-daemon
    #    namespace: openshift-machine-config-operator
      drain:
        # image which provides the openstackclient and has the osc-placement plugin installed
        drainPodImage: quay.io/openstack-k8s-operators/tripleo-deploy
        # enable automatic live migration of instances in ACTIVE state
        # all other instances in different state need to be cleaned up manually
        enabled: false
      openStackClientAdminSecret: openstackclient-admin
      openStackClientConfigMap: openstackclient
      serviceAccount: compute-node

Apply the CR:

    oc apply -f config/samples/compute-node_v1alpha1_computenodeopenstack.yaml
    
    oc get pods -n openstack
    NAME                                   READY   STATUS    RESTARTS   AGE
    compute-node-operator-ffd64796-vshg6   1/1     Running   0          119s

Get the generated machineconfig and machinesets

    oc get machineset -n openshift-machine-api
    oc get machineconfigpool
    oc get machineconfig


## POST steps

### Add compute workers

Edit the computenodeopenstack CR:

    oc -n openstack edit computenodeopenstacks.compute-node.openstack.org example-computenodeopenstack
    # Modify the number of workers and exit

    oc get machineset -n openshift-machine-api
    # check the desired amount has been updated

### Remove compute workers

There are different ways to remove a compute worker node from the OpenStack environment:

#### Remove a random compute worker node

Edit the computenodeopenstack CR and lower the workers number, save and exit:

    oc -n openstack edit computenodeopenstacks.compute-node.openstack.org example-computenodeopenstack

OCP will choose a compute worker node to be removed. Per the `deletePolicy` of the compute worker machineset is set to `Newest`, therefore the newest compute node is expected to be removed. Depending on the status of the instances and `drain` configuration, manual migration/cleanup of the instances is required.

The following command will show which compute worker node got disabled to be removed.

    oc get nodes

#### Remove a specific compute worker node

Edit the computenodeopenstack CR, lower the workers number and add a `nodesToDelete` section in the spec with the details of the worker node which should be removed:

    oc -n openstack edit computenodeopenstacks.compute-node.openstack.org example-computenodeopenstack

Modify the number of `workers` and add the `nodesToDelete` section to the spec:

    nodesToDelete:
    - name: <name of the compute worker node, e.g. worker-3>
      # enable disable live migration for this node. If not specified the global setting from the
      # `drain:` section is used. If not specified there, the default is `false`
      drain: true/false

Save and exit.

## Cleanup

First delete all instances running on the OCP, then delete the operator using OLM.
