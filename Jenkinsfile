node {
  stage 'Build and Test'
  def workspace = pwd()
  env.GOPATH="${workspace}"

  sh 'mkdir {src,pkg,bin}'
  sh 'go get -d -v github.com/superboum/moolinet/...'
  sh 'git --git-dir src/github.com/docker/docker/.git checkout 667315576fac663bd80bbada4364413692e57ac6'
  sh 'go test -v ./...'
}
