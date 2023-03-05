workspace {
	model {
		
		system1 = softwareSystem "System 1" {
			description "This is system 1"
			tags  "web"  "backend" 
		}
		
		system2 = softwareSystem "System 2" {
			description "This is system 2"
			tags  "worker"  "backend" 
		}
		
		
		production = deploymentEnvironment "Production" {
			
			deploymentNode "Web Server" {
				description "Serves the system 1 web application"
				technology "Apache Tomcat"
				
				softwareSystemInstance "System1" {
				}
				
			}
			
			deploymentNode "Worker" {
				description "Processes background jobs"
				technology "Celery"
				
				softwareSystemInstance "System2" {
				}
				
				softwareSystemInstance "System1" {
				}
				
			}
			
		}
	  	
	}

	views {
        deployment * production {
            include *
            autoLayout lr
        }
    }
}