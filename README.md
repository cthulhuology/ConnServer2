ConnServer2
===========

ConnServer2 is a modern reimplementation of a classic MMO game server backend.  The original project
can be found on Github at: https://github.com/cthulhuology/ConnServer and was used in a number of production
games over a number of years.  ConnServer2 is a modern adaptation of the ConnServer codebase building upon
the changes in web programming over the last 8 years.


Changes from ConnServer
-----------------------

Over the past few years a number of changes have occured in the ecosystem of web gaming:

	* Flash has a less dominant position
	* WebSockets mean that web browsers now have full bidirectional messaging
	* HTML5 Canvas and WebGL mean it is possible to build rich UIs in the browser
	* Faster JavaScript engines mean that more processing can be done on the front end
	* JSON has won out over URL encoded strings
	* Databases like PostgreSQL now natively support JSON in key/value stores

The original ConnServer code base was heavily based upon the limitations of using XMLSockets in Flash
to provide full bidirectional real time communication between the browser and the server.  Additionally,
the ConnServer spent much of it's time processing object state on the behalf of the client, as the client
couldn't be trusted to simulate entities and render at 12fps on the client.  As such the original ConnServer
contained an entire distributed in memory cache (similar to memcache meets mongodb) with binary blobs backed
in a Postgres database.

As such the new ConnServer2 code base is going to modernize a bit and attempt to gain a bit of ground over 
it's predicessor.  ConnServer2 will support:

	* JSON as it's native messaging format
	* WebSocket support will work out of the box, as well as raw TCP connections
	* The new Object Server functionality will be migrated to a separate process space
	* An advanced Actor model will be supported for native bots and client connections
	* Multiple routing topologies will be supported, in addition to pub/sub models

Overall ConnServer2 should be slightly more robust and capable of running on substantially more cores than 
the original ConnServer codebase which was started over 11 years ago.  

In addition to the major functional upgrades, I am going to use ConnServer2 as an excuse to evaluate new 
programming languages for system software.  Whereas the original ConnServer was written with an eye towards
making money, ConnServer2 is largely an exercise in improving the original.  Towards that end, I expect to
implement 2 different versions of the same application.

As the original ConnServer was written in C++, the first language that I am going to attempt to implement
it in is Go, http://golang.org.  While the original code base exploited many advanced features of C++, the
Go rewrite will attempt use all of the major features of the language.  This will be the first substantial
program that I am writing in Go, and I'm already finding some of the limitations of the language to be 
a bit more than annoying.  But from Algol in funny hats to Algol in funny hats isn't such a big switch.


Scripting Options and Engines
-----------------------------

Most of the game logic in the original ConnServer based projects were relegated to AI bots that ran on 
separate servers, and were written in scripting languages that would allow for rapid modification to the
working AIs.  However, the original ConnServer also allowed for loading C++ libraries at runtime which could
dynamically intercept and process messages on the server side.  The module system required that each C++
object be versioned, and dynamic code loading required serious games with the symbol tables.

ConnServer2 will attempt to preserve the module behavior as well, however Go currently doesn't support 
dynamically loading packages, and so the support for new modules will be by launching new Go servers and 
migrating the objects from old to new.  While not ideal, it will make it possible to continue to develop 
fast path actors while preserving the ability to keep more complex models out of the critical path.


Roadmap
-------

The following is my thoughts on when I expect to have various bit generally done.  All of these dates are
just hand waving at the moment.  I suspect the basic functionality represented in v0.1 will actually be
enough to build useful systems with, but it will probably take a year to harden and tool it sufficiently well
that one would actually consider using it in production.

Release v0.1 	(Octoberish 2014)
	* messages
	* actors
	* direct routing
	* fanout routing
	* topic routing
	* rooms and announcements
	* actor registry
	* websockets

Release v0.2 	(Novemberish 2014)
	* object cache
	* object persistence

Release v0.3	(Januaryish 2015)
	* live upgrades

Release v0.4	(Marchish 2015)
	* server to server
	* automatic failover

Release v0.5	(Mayish 2015)
	* tcp sockets
	* bot toolkit
	* object toolkit

Release v0.6	(Julyish 2015)
	* ui toolkit
	* management ui

Release v0.7	(Septemberish 2015)
	* sample game

Release v1.0	(Octoberish 2015)
	* fully operational deathstar


NB: These dates all assume that life doesn't find ways to prevent me from hacking on this occasionally.


Contact & Questions
-------------------

If you are interested in the development of ConnServer2, or would like to use it for some purpose I never
intended, please feel free to drop me a line at dave at dloh.org.


License
-------

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero Public License for more details.

    You should have received a copy of the GNU Affero Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
