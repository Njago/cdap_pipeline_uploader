# cdap_pipeline_uploader


## Function
This will loop through a directory of .Json files exported from cdap and upload all files in that directory to the specified namespace.

### Use


| Variable  	| Use                                                               
|-----------	|-------------------------------------------------------------------
| dir       	| Directory to pull files from                                      
| namespace 	| Namespace to upload pipelines to                                  
| host      	| Default is for cdap. Change to CDF instance url/api for CDF use 
| authToken     | Authtoken from gcloud auth for CDF use
| trim          | Default is "-cdap-data-pipeline.json" for cdap
