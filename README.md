# cdap_pipeline_uploader


## Function
This will loop through a directory of .Json files exported from cdap and upload all files in that directory to the specified namespace.

### Use


| Variable  	| Use                                                               
|-----------	|-------------------------------------------------------------------
| dir       	| directory to pull files from                                      
| namespace 	| namespace to upload pipelines to                                  
| host      	| default is for cdap. change to <CDF instance url>/api for CDF use 
| authToken     | authtoken from gcloud auth for CDF use
| trim          | default is -cdap-data-pipeline.json for cdap
