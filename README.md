<html>
<body>
<h1 class="title">DBA API adapter</h1>
<h3 class="why">Why</h3>
<p class="why">This project was created because I needed to buy a bed and wanted to learn Go.</p>
<h3 class="what">What</h3>
<p class="what">DBA (Den blå avis) is the largest danish intermediary portal for buying and selling between private individuals. It does not expose a public API though, so it makes it difficult to create anything for tracking new offers etc. I decided to create a simple web scraper that converts the query data into a JSON object, that can be accessed via a Web API.</p>
<h3 class="how">How</h3>
<p class="how">A simple Web API was written in Go, usin Gin HTTP web framework and Colly - Go scraper. A single request makes the server scrap all the pages related to the query on DBA, and return the extracted data as a JSON. Also a Client using the API was created, but it's not a part of this repo and is not publicly available.</p>
<h3 class="usage">How to use</h3>
<p class="usage">You need to have docker installed. Once that's done add an entry mapping your loopback address to <i>scraper.docker</i> to your host file (or use any of the available proxy servers to do that). Then <i>docker-compose up</i> the hell out of it. The root url is <i>scraper.docker/api</i>, simply append the path you'd use within DBA website and get a beautiful JSON returned. Example:
you can get the data from <i>https://dba.dk/biler/biler</i> at <i>http://scraper.docker/api/biler/biler</i>. Enjoy!</p>
<h3 class="technologies">Technologies used</h3>
<ul class="technologies">
  <li class="technologies" hover="Python">Go</li>
  <li class="technologies" hover="Go HTTP web framework">Gin</li>
  <li class="technologies" hover="Go web scraper implementation">Colly</li>
</ul>
<hr>
<small class="created">Created: August 2020</small>
</body>
</html>
