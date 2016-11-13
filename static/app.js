angular.module('moolinet', ['ngResource', 'ngRoute'])

  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/challenge', {
        templateUrl: 'partials/challenge.html',
        controller: 'ChallengeListController'
      })
      .when('/challenge/:slug', {
        templateUrl: 'partials/challenge_view.html',
        controller: 'ChallengeViewController'
      })
      .when('/hall-of-fame', {
      })
      .when('/configuration', {
      })
      .when('/login', {
      })
      .otherwise({
        redirectTo: '/challenge'
      })
  }])

  .factory('Challenge', ['$resource', function($resource) {
    return $resource('/api/challenge/:slug', {slug:'@id'});
  }])

  .controller('ChallengeListController', ['Challenge', function(Challenge) {
    this.list = Challenge.query();
  }])

  .controller('ChallengeViewController', ['Challenge', function(Challenge) {
  }])
;
