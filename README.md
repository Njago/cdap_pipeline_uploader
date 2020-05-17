# cdap_pipeline_uploader


## Function
This will loop through a directory of .Json files exported from cdap and upload all files in that directory to the specified namespace.

### Use

set dir - directory to pull files from
set namespace - namespace to upload to
set host - change to instance url/api for CDF use
set authtoken - authtoken from gcloud auth for CDF use
set trim - default is -cdap-data-pipeline.json for cdap