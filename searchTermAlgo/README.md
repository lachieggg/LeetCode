# SearchTermAlgorithm

History
=

- This is an old program, originally authored in 2016.

- It is here, at this stage, for historical purposes only.

Notes
=
    
This program was written for [COMP10002](https://handbook.unimelb.edu.au/subjects/comp10002/).

This program, given a text file of arbitrary size,
searches for the line which resembles the search term
as closely as possible. It will return a rank
of the lines in the text which achieved the highest level 
of similarity.

Resemblance is defined by the number of consecutively
correct characters a line has with respect to a query.

However, the part of the line which matches the query
could be a subset of the query, in fact it happens to be
that even for large texts, queries of more than a few words
often will not be matched. This is the advantage that this program
has over a binary "yes this is a match" or "no it is not" 
algorithm.

In short, this program will check all possible consecutive subsets of 
the line against all possible subsets of our query. 

Future
=
My goal with this program is to implement more complexity to `get_score()`.

Allowing us to check a variable number of parts of the line 
for subsets of our query.

License
=
    Copyright (C) 2016 Lachlan Grant
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
    
