
## Raumschiff

![alt text](img/teaser.PNG "1")

**Das folgende Kapitel befasst sich mit dem Einsatz von sog. "Modifier. Behandelt werden der Mirror-Modifier und der Boolean-Modifier**



## 1. Vorbereitung

* Belasst eure Szene wie sie ist.

* Sucht euch Referenzbilder zu einem Raumschiff eurer Wahl

* Fügt ein Background-Image ein. (Siehe Raumhafen)



## 2. Modifier anwenden

![alt text](img/1_modifier.PNG "1")

* Löscht den Würfel insofern dies für euer Raumschiff nötig sein wird, achtet beim einfügen des neuen Raumschiffs darauf, dass es sich auf den 0-Korrdinaten befindet.
* Wählt euer Grundobjekt aus. In den Reitern unter eurer Szenen-Hierarchie findet ihr einen Schraubschlüssel. Wählt ihn aus und wählt anschließend "Mirror".

## 3. Modellierung


![alt text](img/1_modeling_mirror.PNG "1")

* Wechselt mit dem ausgewählten Objekt in den Edit-Mode und verschiebt es etwas auf der roten X-Achse. Euch wird auffallen, dass das Objekt auf den anderen Seite gespiegelt wird. 

* Aktiviert im Mirror-Modifier "Clipping", damit eure Modelle nicht über den Spiegelungspunkt hinaus geschoben werden können. Schiebt es nun in kontakt mit dem gespiegelten Würfel.

* In diesert Übung werden wir einen A-Wing aus dem Star Wars Universum modellieren. Hierfür eignet sich ein Würfel als Grundform.

![alt text](img/3_modeling_shape.PNG "1")

* Der Würfel muss für die Grundform des A-Wing an der Spitze flah und weniger breit und am hinteren Teil Hoch und breit sein. Wendet nun Skalierungen und Translationen an um die gewünshte Form zu erhalten.






## 4. Rumpf

![alt text](img/4_shaping.PNG "1")

* Für eine bessere Übersicht empfiehlt es sich, mit `Z` in die "Wire-Frame-View" zu wechseln un beim weiteren vorgehen mit `Num 7`in die "Top-View" zu wechseln.

* Schneidet der länge nach mit `Ctrl+R` und `LMB` einen "Edge-Loop" in den Rumpf des Schiffs.

![alt text](img/4_shaping_cuts.PNG "1")


* wechselt in den "Face-Mode" mit `Ctrl+Tab`, wählt das neu erzeugte "Face" und extrudiert es entsprechend in Richtung Heck.

* Schneidet nun horizontal durch den Rumpf, um die Abflachung an den Lasergeschützen zu formen. Wechselt nun mit `Ctrl+Tab` in den "Vertice-Mode", selektiert alle 4 der Vertices, an denen sich die Flache Stelle befinden soll und skaliert sie entlang der X-Achse mit `S` gefolgt von `X` bis sie flach ist.

* Ein weiterer Horizontaler Schnitt ist an der Schnauze Notwendig. Geht vor wie in den vorherigen Schritten beschrieben.

## 5. Hard Surface Modeling

![alt text](img/raumschiff_hardSurface1.PNG)

* Um das Aussehen einer harten Kante zu erzielen, ist es vorallem bei Benutzung von der "Subdivision Surface Modifier", "Multiresoltuion Modifier" oder von "Shade Smooth" nötig, diese Kanten mit einer "Crease" Operation zu bearbeiten, damit sie nicht abgerundet werden.

* Im folgenden Beispiel wurde der "Mirror Modifier" durch drücken des "Apply" Buttons angewandt (es wird empfohlen das mittlere "Face" zwischen den gespiegelten Hälften zuvor zu löschen).

* Zuerst fügt im einen "Subdivision Surface Modifer" hinzu. Wechselt in den "Edit Mode" mit `Tab`.

* Wählt anschließend alle Kanten aus (auch die vertikalen) die Hart erscheinen sollen. Gebt nun im rechten Menü unter "Mean Crease" (oder Crease, sollte ihr nur eine "Edge" gewählt haben) die zahl 1 ein.

![alt text](img/raumschiff_hardSurface2.PNG)

* Wenn ihr nun mit `Tab` den "Edit Mode" verlasst seht ihr, dass die zuvor abgerundeten Kanten nun harte Kanten sind.

## 6. Schubdüsen

![alt text](img/5_jets_cockpit.PNG "1")

* Ein weiteres drücken von `Z` bringt euch zurück in die "Solid-View".

* Wechselt mit `Tab` zurück in den Objekt-Mode und fügt mit `Shift+A` einen Cylinder ein. Sorgt mit F6 dafür, dass er nicht zu viele Polygone besitzt.

* Schiebt ihn an die richtige Position mittels `G` gefolgt von der Achse auf der ihr verschieben wollt (`X`,`Y` oder `Z`). Rotiert den Zylinder entlang der X-Achse in die entsprechende Richtung.

* Ändert die Form des Zylinders so, dass sie am vorderen Teil dünner auisfällt als am hinteren. Es werden ggf. zusätzliche Edge-Loops benötigt, um die Form anzupassen.


## 7. Cockpit

![alt text](img/6_cockpit.PNG "1")

* Das Vorgehen ist nun das selbe wie beim Rumpf. Erzeugt einen Würfel, fügt einen "Mirror-Modifier" hinzu und so weiter.

* Für das Cockpit müsst ihr mindestens zwei weitere "Edge-Loops" hinzufügen um die Form herauszuarbeiten.

* Nutzt die Grundobperationen (Translation, Rotation und Skalierung) um die Form im "Edit-Mode" zu erzeugen.

![alt text](img/6_cockpit_merged.PNG "1")

* Wendet erst den Mirror-an. Fügt anschließend das Cockpit eurem Rumpf hinzu. Selektiert zuerst euer Cockpit und anschließend mit `Shift+RMB` den Rumpf. Drückt anschließend `Ctrl+J` um die Objekte zu mergen.

![alt text](img/6_cockpit_steering.PNG "1")

* Fügt für die Steuerungsklappen zwei Aufrechte, schmal-skalierte Quader ein. und verbindet die beiden Meshes mit `Alt+J`.

---
date: 2016-03-09T00:11:02+01:00
title: Übung 4.1 - Raumschiff
weight: 41
---


* Fügt, nach belieben einen Toon-Shader ein.

**Ihr könnt entweder bereits bei den Düsen oder erst beim Cockpit die Vertikalen Steuerungsklappen hinzufügen, mit Hilfe des Boolean-Modifiers können mit der Auswahl von Difference die Aushöhlungen in die Düsen geschnitten werden.**





## Aufgabe

* Erstellt einen X-Wing oder ein anderes Schiff, das sich von Grundauf von eurem bereits erstellten unterscheidet.



## Freiwillige Zusatzaufgabe

![alt text](img/additional.PNG "1")

* Überarbeitet beide Schiffe so, dass sie mehr detailles erhalten und ihren Vorbildern ähnlicher werden.



## Gelerntes

Aktion                 | Keyboard-Shortcut                  | Menübefehl 
-----------------------|------------------------------------|------------
Mirror-Modifer ||
Edge-Loops hinzufügen |`Ctrl+R+LMB`| Selected Meshes -> Object -> Join
Hard Surface Modeling | |