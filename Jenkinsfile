node {
  stage('Configure') {
    deleteDir()
    def workspace = pwd()
    env.GOPATH="${workspace}"
    env.PATH=env.PATH+":"+env.GOPATH+"/bin"
    sh 'mkdir -p bin pkg src src/github.com/superboum/moolinet'
  }

  stage('Build') {
    dir('src/github.com/superboum/moolinet') {
        checkout scm
        sh 'make prepare docker=1.12'
        sh 'make'
        archiveArtifacts 'release/**/*'
    }
  }

  stage('Test') {
    dir('src/github.com/superboum/moolinet') {
        sh 'make test'
    }
  }

  stage('Lint') {
    dir('src/github.com/superboum/moolinet') {
      sh 'make lint'
    }
  }
}

