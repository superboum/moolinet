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
    var Challenge = $resource('/api/challenge/:slug', {slug:'@id'});
    var list = []

    return {
      load: function(cb) { if (list.length == 0) { list = Challenge.query(cb) } else { cb(); } },
      getList: function() { return list; },
      getChallenge: function(slug) { return list.find(function(elem) { return elem.Slug == slug }); }
    }
  }])

  .controller('ChallengeListController', ['Challenge', function(Challenge) {
    Challenge.load(function() { this.list = Challenge.getList(); }.bind(this));
  }])

  .controller('ChallengeViewController', ['Challenge', '$routeParams', function(Challenge, $routeParams) {
    Challenge.load(function() { this.selected = Challenge.getChallenge($routeParams.slug); }.bind(this));
  }])
;
