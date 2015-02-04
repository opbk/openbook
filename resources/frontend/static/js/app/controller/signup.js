'use strict';

(function () {
	var SignUpController = function($scope) {
		$scope.form = {
			email: "",
			password: "",
			new: "false",
			error: {}
		};

		$scope.submit = function($event) {
			var invalid = false
			if($scope.form.email == "") {
				$scope.form.error.email = "Укажите ваш email";
				invalid = true;
			}

			if($scope.form.new == "false" && $scope.form.password == "") {
				$scope.form.error.password = "Укажите ваш пароль";
				invalid = true;
			}

			if(invalid) {
				event.preventDefault();
			}
		}
	};

	angular.module('App').controller(
		'SignUpController', [
			'$scope',
			SignUpController
		]
	);
}());