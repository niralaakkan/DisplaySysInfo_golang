pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'docker build . -t golang'
        echo 'Build complete'
      }
    }

    stage('') {
      steps {
        input(message: 'Input', ok: 'Ok')
      }
    }

  }
}