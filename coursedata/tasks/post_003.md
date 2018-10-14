
## Polygon Modeling

![alt text](img/teaser.PNG "1")

**Das folgende Kapitel befasst sich mit der Erstellung eines Gesichts durch die Verwendung von Polygon Modeling unter der Berücksichtigung des Edge-Flow**



## 1. Vorbereitung

![alt text](img/1_vorbereitung.PNG)

* Löscht eure Startszene und fügt statt des Quaders eine Plane ein.

* `Ctrl+A`, dann `X` gefolgt von `Shift+A`-> Plane.

* Rotiert die Plane mit `R` gefolgt von `X` und `Num 90` um 90° auf der X-Achse.

* Fügt einen Mirror-Modifier hinzu insofern eurer gewähltes Gesicht symmetrisch ist.



## 2. Edge-Flow eines Gesichts

![alt text](img/edgeFlow.PNG )

* Das Bild deutet mit Linien an, wie der grobe Edge-Flow der hälfte des Gesichts aussehen sollte.

* Setzt als Hintergrundbild ein Gesicht (von Jabba), wie im Kapitel 
[Raumhafen](https://sftp.hs-furtwangen.de/~mch/computergrafik/script/chapter01/exercise01/#2-background-image) erklärt wurde. Iealerweise findet ihr ein Bild der Front und ein Bild in der Seitenansicht des selben Gesichts.

* Wählt eure Plane aus und wechselt mit `Tab` in den Edit-Mode.



## 3. Grundlegende Edges erzeugen

![alt text](img/grundlegende_edges.PNG)

* Löscht eine Hälfte des Planes mit `X`, so dass nurnoch eine Edge mit zwei Vertices übrig bleibt.

* Stellt sicher, dass ihr in der Orthogonalen Frontansicht seid (`Num 5` gefolgt von `Num 1`).

![alt text](img/grundlegende_edges2.PNG)

* Wählt euch eine der in 2. gezeigten Edge-Flow-Linien und beginnt in der Frontansicht die Form, angepasst auf euer Gesicht komplett nachzumodellieren, indem ihr die beiden Vertices der Plane mit `E` extrudiert, Faces aus diesen erzeugt und schließlich den ersten Edge-Loop durch Extrudieren der Face-Kanten erstellt. Bei Jabba ist ein Edge-Loop für die Nasenspitze nicht nötig und wurde im Beispiel ausgelassen, weil er keine Nasenspitze besitzt. 

![alt text](img/face_bridgeLoop.PNG)

* Wie im Bild zu sehen, wurde darauf geachtet, dass der ausgewählte Loop gleich viele Vertices besitzt wie der , mit der er durch den Bridge Befehl verbunden werden soll. Eine gleiche Verticeanzahl ist dringend notwendig, da der Bridge Befehl ansonsten Tri-Faces oder im schlimmsten Fall ein Polygonchaos mit unübersichtlichen Verbindungen erzeugt. Achtet also schon beim erstellen der ersten Edge-Loops darauf, dass spätere Edgeloops die gleiche Verticeanzahl besitzen. Achtet außerdem darauf, dass ihr keine Tris (Faces mit drei Vertices) erzeugt, da diese den u.a. Edge-Flow unterbrechen.

* Zurück zu Jabba: Wiederholt den Bridge-Vorgang für einen jeden weiteren Loop.

> Beim hinzufügen von neuen Polygonreihen, extrudiert nicht einfach die Polygone zu einer neuen Reihe. Dabei entstehen Löcher im Mesh. Stattdessen extrudiert zuerst die Punkte und füllt dann die Polygone auf.
![](img/v1.gif)
![](img/v2.gif)




* Sobald ihr alle Edges erstellt habt, geht es weiter zur Tiefe des Gesichts.



## 4. Tiefe

![alt text](img/tiefe.PNG)

* Da euer Modell im Moment nur Zweidimenstional ist, wird es Zeit Jabbas Gesicht etwas 3D zu verleihen.

* Bei diesem Schritt wird ein häufiger Wechsel zwischen Front-(`Num 1`) und Seitenansicht (`Num 3` oder `Num+Alt 3`) benötigt.

* Markiert die benötigten Edges, Faces, oder Vertices (`Rmb` oder `Alt+Rmb` für Edge-Loops) und zieht sie mit `G` gefolgt von `Y` entsprechend in die Tiefe.


## 5. Edges-Verbinden

![alt text](img/edges_verbinden.PNG)

* Als letzten Schritt verbindet ihr die einzelnen Edge-Loops. Markiert sich gegenüberliegende Vertices und fügt ein Face zwischen ihnen ein, indem ihr `F` drückt. Ihr könnt auch viele Vertices miteinander verbinden indem ihr den Bridge Edge-Loops Befehl verwendet. Ihr findet diesen indem ihr die `Space` drückt und anschließend den Befehlsnamen eingebt. Damit dieser sauber funktioniert muss die gleiche Anzahl von Vertices von der einen un von der anderen Seite gewählt sein.

* Nachdem das Gesicht komplett verbunden wurde, könnt ihr mit `Ctrl+R` neue Edge-Loops einfügen um einen höheren Detailgrad zu erhalten.

![alt text](img/finished.PNG "1")


## Aufgabe

* Jabba-Gesicht fertig stellen.


## Freiwillige Zusatzaufgaben

* Vernäht das Polygon Gesicht mit euren erzeugten Jabba-Grundkörper. 
* Erzeugt ein (zusätzliches) humanoides Gesicht.


## Gelerntes

Aktion                 | Keyboard-Shortcut                  | Menübefehl 
-----------------------|------------------------------------|------------
Edge-Cut    | `Ctrl+R` |
Polygon Modeling           |                         | 
Extrudieren          |             `E`            | Mesh -> Extrude
Rotate          |                    `R`     | Mesh -> Transform -> Rotate
Translate           |        `G`                 | Mesh -> Transform -> Grab/Move
Bridge Edge Loop           |                      | Mesh -> Edges -> Bridge Edge Loops


