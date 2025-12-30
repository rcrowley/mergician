TODO
====

Digital gardening potential
---------------------------

Mergician, Electrostatic, and Feed trivially allow folks to put content wherever they want, date what content they want, and make a feed of what dated content they want. This makes it ideal for a digital gardening style of site. If the author wants, nothing about the site but the Atom feed has to be in reverse-chronological order.

Once it exists, Dirindex will further the duality of hierarchical directory structure for content with a reverse-chronological feed built from some of it.

Dirindex
--------

Simple. Dirindex should automatically create an `index.html` in any directory that's missing one and which contains files in it or any subdirectory.

It's possible this should just be an option to Electrostatic. That'd make merging into the right site chrome easy but make it hard to apply selectively.

Bidirectional links
-------------------

Bidirectional links a la Roam Research or Notion are back in style. I do enough self-linking that there would definitely be utility, though I'm not passionate about figuring out how to do this.

Allowing bidirectional links from other sites like blog pingbacks/trackbacks/webmentions would require reading `Referer` headers from logs or implementing one of those old standards, all of which seems like a waste of time, prone to spam, and in need of human-in-the-loop acceptance/rejection.

Whatever the source, styling the bidirectional links as a grey section below even the footer, almost like debug information, sounds aesthetically pleasing.
