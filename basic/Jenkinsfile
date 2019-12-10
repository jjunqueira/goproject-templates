pipeline {
    agent any
    tools {
        go 'Go 1.13'
    }
    stages {

        stage ('Test') {
            steps{
                sshagent (credentials: ['golang-ssh-key']) {
                    sh 'make clean'
                    sh 'make test'
                }
            }
        }
      
    	stage ('Build') {
            steps{
                sshagent (credentials: ['golang-ssh-key']) {
                    sh 'make release'
                }
                archiveArtifacts allowEmptyArchive: false, artifacts: 'target/*.gz', caseSensitive: true, defaultExcludes: true, fingerprint: false, onlyIfSuccessful: false 
            }
    	}
    }
}