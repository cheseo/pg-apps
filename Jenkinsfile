pipeline {
	agent any
	stages {
		stage ('all') {
		dir("/home/public/cloud-scripts/week-7/jenkins-terra"){
			steps {
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
