'use strict';

(function () {
	var HistoryController = function($scope, $http, toaster) {
		$scope.order_hide = {}
		$scope.book_show = {}

		$scope.delete = function(id) {
			if(confirm('Вы дествительно хотите удалить заказ?')) {
				$http.delete('/order/' + id)
					.success(function(data, status, headers, config) {
						console.log("order deleted successefully")
						$scope.order_hide[id] = true;
						toaster.pop('success', '', 'Заказ успешно удален')
					})
					.error(function(data, status, headers, config) {
						toaster.pop('error', '', 'При удалении заказа возникла ошибка')
					})
			}
		}
	};

	angular.module('App').controller(
		'HistoryController', [
			'$scope',
			'$http',
			'toaster',
			HistoryController
		]
	);
}());