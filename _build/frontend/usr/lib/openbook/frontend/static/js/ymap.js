ymaps.ready(function() {
	var myMap;
	$('.fancybox').fancybox({height:600, afterShow : function() {
		myMap = new ymaps.Map('map-container', {
		    center: [55.752037,37.616792],
		    zoom: 14,
			behaviors: ["scrollZoom","drag"]
		});	
 
		var myPlacemark = new ymaps.Placemark([55.752037,37.616792], {
	        iconContent: 'Офис самовывоза "Открытая книга"',
	        balloonContentHeader: 'Офис самовывоза "Открытая книга"',
	        balloonContentBody: 'Тут будет наш адрес',
	        balloonContentFooter: 'Сайт: <a rel="nofollow" href="http://opbook.org/" target="_blank">http://opbook.org</a>'
	    }, {
	        preset: 'twirl#redStretchyIcon',
			balloonMaxWidth: 250
	    });

		myMap.geoObjects.add(myPlacemark);			
 
	}, afterClose:function (){
	    myMap.destroy();		
		myMap = null;
	}});
});
