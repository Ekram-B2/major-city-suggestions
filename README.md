# Suggestions Manager

A small go application built designed to support the core functionality of being
able to support a retreival of city suggestions given a search term.

The application is built in mind of being maintainble in the sense of being easy
to makes changes to, and minimal in regards to what the composing components are
to do. 


The service is availible at the end point
https://immense-journey-27252.herokuapp.com/

The application is meant to be paired with the suggestions cache application which supports caching of web requests
made to the worker process which computes the rank, and the rank manager that computes actual rank

https://github.com/Ekram-B2/suggestionscache
https://github.com/Ekram-B2/rankmanager


Some of the input information was pre-processed prior to its use within the application. The notebook where this occured is found here:

https://github.com/Ekram-B2/city-analyzer


