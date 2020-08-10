<html>
<body>
<h1 class="title">DBA API adapter</h1>
<h3 class="why">Powód</h3>
<p class="why">Projekt powstał ponieważ potrzebowałem kupić łóżko. A przy okazji spróbować też Go.</p>
<h3 class="what">Cel</h3>
<p class="what">DBA (Den blå avis) jest największym duńskim portalem to sprzedawania i kupowania rzeczy przez między osobami prywatnymi (taki polski OLX). Niestety nie udostępnia on publicznego API, które ułatwiałoby tworzenie aplikacji automatyzujących korzystanie z portalu.</p>
<h3 class="how">Wykonane</h3>
<p class="how">Proste API webowe wykorzystujące Go, Gin - framework HTTP dla Go, oraz Colly - web scraper w Go. Wysłanie requesta do API skutkuje zescrapowaniem stron odpowiadajacych danemu query na DBA i zwróceniu danych w postaci JSONa.</p>
<h3 class="usage">Jak korzystać</h3>
<p class="usage">Wymagany jest Docker. Następnie koniecznie jest zmapowanie adresu loopback komputera do adresu <i>scraper.docker</i> (wprowadz zmiany w pliku hosts lub uzyj servera proxy). Następnie wykonaj <i>docker-compose up</i> w root folderze. Root url to <i>scraper.docker/api</i>, do którego należy dodać ścieżkę dokładnie taką, jaką użylibyśmy na DBA, by otrzymać piękny, sparsowany JSON. Przykład:
aby dostać dane ze strony <i>https://dba.dk/biler/biler</i> wyślij zapytanie GET na adres <i>http://scraper.docker/api/biler/biler</i>. Miłej zabawy!</p>
<h3 class="technologies">Wykorzystane technologie</h3>
<ul class="technologies">
  <li class="technologies" hover="Python">Go</li>
  <li class="technologies" hover="Go HTTP web framework">Gin</li>
  <li class="technologies" hover="Go web scraper">Colly</li>
</ul>
<hr>
<small class="created">Data powstania: Wrzesień 2020</small>
</body>
</html>
