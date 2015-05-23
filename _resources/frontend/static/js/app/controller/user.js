'use strict';

(function () {
	var UserController = function($scope, Address) {
		$scope.form = {
			name: "",
			phone: "",
			error: {}
		};

		$scope.addAddress = function($event) {
			Address.save($scope.address, function(address) {
				$scope.addresses.push(address);
				$scope.form.address = address.id;
				$scope.addAddressFlag = false;
				$scope.address = {};
			}, function(error) {
				alert('Произовшла ошибка при добавление адреса. Попробуйте еще раз через несколько минут.');
			});
		}

		$scope.submit = function($event) {
			var invalid = false;
			if($scope.form.name == "") {
				$scope.form.error.name = "Укажите как к Вам обращаться";
				invalid = true;
			}

			if($scope.form.phone == "") {
				$scope.form.error.phone = "Укажите телефон, что бы мы могли связаться с Вами";
				invalid = true;
			}

			if($scope.form.password != "" && $scope.form.rpassword) {
				$scope.form.error.password = "Введенные пароли не совпадают";
				invalid = true;
			}

			if(invalid) {
				event.preventDefault();
			}
		}
	};

	angular.module('App').controller(
		'UserController', [
			'$scope',
			'Address',
			UserController
		]
	);
}());