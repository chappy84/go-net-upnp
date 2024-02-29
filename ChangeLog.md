* 20xx-xx-xx v1.0.0
	* Support UPnP v2.0 specifications more correctly

* 20xx-xx-xx v0.9.0
	* Support the event subscription function of UPnP and deprecated functions from UPnP v1.1 such as query function

* 2024-02-29 v0.8.2
	* Thanks for Tom Chapman (@chappy84)
	* Ensure URL paths don't start with multiple consecutive slashes
    * Can cause some UPnP devices to return a 404 where otherwise valid
    * Correct spelling mistake, make old function name alias for backwards compatibility

* 2022-09-25 v0.8.1
	* Update deprecated functions
	* Add go.mod
	* Fix lint errors

* 2015-07-21 v0.8.0
	* The first public release
	* Support almost all control point and device functions of UPnP v1.0 specification