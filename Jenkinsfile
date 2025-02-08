pipeline {
    agent { docker { image 'golang:1.23.6-alpine3.21' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'go build main.go'
                archiveArtifacts 'main'
            }
        }
    }
}
