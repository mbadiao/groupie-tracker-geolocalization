window.onload = function () {
  let locationstr = document.getElementsByClassName("Location")
  let AllLocations = []

  for (var i = 0; i < locationstr.length; i++) {
    var content = locationstr[i].innerHTML; // ou elements[i].textContent
    AllLocations.push(content)
  }
  console.log(AllLocations)

  L.mapquest.key = 'ffijQuOZJR2ep7lH42nJwF5Iy7KAX8cB';

  // Geocode three locations, then call the createMap callback
  L.mapquest.geocoding().geocode((AllLocations), createMap);

  function createMap(error, response) {
    // Initialize the Map
    var map = L.mapquest.map('map', {
      layers: L.mapquest.tileLayer('map'),
      center: [0, 0],
      zoom: 0,
      maxZoom: 13,
    });

    let newTranche = AllLocations.slice(1, -1)

    map.addControl(L.mapquest.control());

    L.mapquest.directions().route({
      start: AllLocations[0],
      end: AllLocations[AllLocations.length - 1],
      waypoints: newTranche
    });

    // Generate the feature group containing markers from the geocoded locations
    var featureGroup = generateMarkersFeatureGroup(response);

    // Add markers to the map and zoom to the features
    featureGroup.addTo(map);
    map.fitBounds(featureGroup.getBounds());
  }

  function generateMarkersFeatureGroup(response) {
    var group = [];
    for (var i = 0; i < response.results.length; i++) {
      var location = response.results[i].locations[0];
      var locationLatLng = location.latLng;

      // Create a marker for each location
      var marker = L.marker(locationLatLng, {
        icon: L.mapquest.icons.marker()
      }).bindPopup(location.adminArea1 + ', ' + location.adminArea3 + ', ' + location.adminArea5);
      group.push(marker);
    }
    return L.featureGroup(group);
  }
}
