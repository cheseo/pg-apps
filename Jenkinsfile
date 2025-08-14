pipeline {
	agent any
	stages {
		stage ('all') {
		steps {
			dir("/home/public/cloud-scripts/week-7/jenkins-terra"){
			sh '''
			git pull
			terraform init
			terraform plan
			'''
			input message: "Continue?"
			sh '''
			terraform apply -auto-approve
			'''
			}
		}
		}
	}
}
