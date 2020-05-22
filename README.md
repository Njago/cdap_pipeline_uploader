
# cdap_pipeline_uploader

## Function

This will loop through a directory of .Json files exported from cdap and upload all files in that directory to the specified namespace.

### Use

| Variable      | Use
|-----------    |-------------------------------------------------------------------
| dir           | ${DIR} - Env Variable for the Directory taht contains your JSON files
| namespace     | ${NAMESPACE} - Namespace to upload to
| host          | ${CDAP_ENDPOINT} - Endpoint for PUT command. 
| authToken     | ${AUTH_TOKEN} - GCloud Auth Token for CDF

### Terminal commands to fill CDF Enviornment Variables

gcloud auth login
export AUTH_TOKEN=$(gcloud auth print-access-token)
export LOCATION=region
export INSTANCE_ID=instance-id
export CDAP_ENDPOINT=$(gcloud beta data-fusion instances describe \
    --location=${LOCATION} \
    --format="value(apiEndpoint)" \
  ${INSTANCE_ID})