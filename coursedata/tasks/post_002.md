## Astromech Droide
![alt text](img/start.PNG "Lesson's Result")

**Damit der Raumhafen nicht so leer aussieht erstellen wir jetzt einige Astromech Droide, die sich um die künftigen Raumschiffe kümmern.**
**Die folgende Übung befasst sich mit grundlegenden Funktionen wie Rotation und Translation sowie Ansichten in Blender. Außerdem werden einige Grundtechniken des Arbeitens im Edit-Mode vorgestellt**


## 1. Vorbereitung 
![alt text](img/1_selectAll.PNG "Selet All")

* Startet Blender. Im neuen Projekt drückt ihr `A`, um alles zu Markieren und anschließend `X` gefolgt von einem klick auf "Delete" um eure Szene komplett zu bereinigen.

* Drückt nun `Num 5`, um von der Perspektiven-Darstellung in die Orthogonal-Darstellung zu wechseln. Drückt anschließend `Num 1`, um auf die Frontansicht zu wechseln.


## 2. Torso
![alt text](img/2_torsoCylinder.PNG "Insert Mesh")

Zunächst brauchen wir eine Grundfrom für den Torso. Ein Zylinder ist perfekt.

* Um einen Mesh einzufügen drückt ihr `Shift+A`. Wählt nun, unter dem Reiter "Mesh" den "Cylinder" aus. 

* Bevor ihr den Zylinder bewegt o.ä. könnt ihr die Polygone reduzieren, indem ihr `F6` drückt und bei Vertices einen niedrigeren Wert eintippt.

* Achtung: Je niedriger der Wert, desto kantiger wird der Torso. Je Höher der Wert, desto runder wird der Torso, hat aber schnell unnötig viele Polygone und erschwert die kontrollierte Weiterarbeit.


## 3. Kopf
![alt text](img/3_head.PNG "Head")

Bevor wir zum Kopf gelangen, benennen wir den Zylinder des Torsos zunächst in Torso um. Rechts oben befindet sich die Objekt-Hierarchie. Doppelklickt auf den Namen "Cylinder" und ändert ihn in Torso.
* Diesen Schritt wiederholt ihr im späteren Verlauf für jedes neue Objekt und weist einen passenden Namen zu.

* Der Kopf soll die selbe Grundgröße haben wie der Zylinder. Wählt den Zylinder im Object-Mode aus und wechselt mit `Tab` in den Edit-Mode. Drückt dann `Ctrl+Tab`, wählt dort den Face-Mode und selektiert anschließend das oberste Face. Drückt dann `Shift+D` um ein Duplikat anzufertigen. Nun löst ihr das duplizierte Face: mit einem Druck auf `P` wählt nun von Selection. Damit wird ein neues Objekt erzeugt, das nur aus der kopierten oberen Deckfläche des Zylinders besteht.

![alt text](img/3_head_1.PNG "Head")

* Wechsel zurück in den Object-Mode und selektiert euer abgetrenntes Face. Wechselt nun wieder in den Edit-Mode, wählt das Face aus und extrudiert entlang der Z-Achse indem ihr `E` gefolgt von `Z` drückt. Zieht das extrudierte Face nun soweit nach oben, bis euch die Kopfhöhe genügt. Skaliert den das oberste Face abschließend mit `S` und einer Mausbewegung kleiner. Bestätigt mit `LMB`.

![alt text](img/3_head_2.PNG "Head")

* Ein doppeltes drücken von `Z` würde eine Skalierung entlang der lokalen Koordinaten des Objekt bewirken, das benötigen wir für den Moment allerdings nicht.



## 4. Beine
![alt text](img/5_leg.PNG "Legs")

Für die Beine benötigen wir zwei Cubes.

* Fügt einen Cube-Mesh ein. und schiebt ihn am roten Pfeil etwas zur Seite. Skaliert ihn nun erst entlang der Y-Achse und der X-Achse. Für die Y-Achsen Skalierung bietet es sich an, die Ansicht auf dei Seitenansicht zu schalten. 
Drückt hierfür `Num 3`. 

![alt text](img/5_legMore.PNG "Leg Advanced")

* Falls ihr euer Objekt zuvor nach Links verschoben habt, könnt ihr auf die entgegengesetzte Seite schauen indem ihr `Ctrl+Num 3` drückt.

* Ihr könnt die Ansichten, solltet ihr mit den Tastenkürzeln nicht zurecht kommen auch über das "View" im Menü unter eurer 3D-Ansicht wechseln.



## 6. Beinform
![alt text](img/6_legShape.PNG "Shaping the Leg")

* Wählt das Bein und wechselt mit `Tab` in den Edit-Mode. Damit wir nun Platz für die Rollen schaffen können, müssen wir dem Bein im unteren Bereich einen Edge-Loop hinzufügen. Drückt dafür `Ctrl+R` und fahrt mit der Maus über eine der vertikalen Linien. Ein Klick sorgt für das Platzieren des Edge-Loops. 
* Des weiteren könnt ihr mit der Maus die Höhe des Schnitts variieren. `LMB` bestätigt die Position.
* Nun wollen wir die unterste Kante selektieren. Drückt während ihr im Edit Mode seid `Z` um in den Wireframe Mode zu wechseln. Drückt nun `B` und zieht eine Selektionsbox über die unterste Reihe Vertices auf.
* Sobald ihr die Vertices selektiert habt, skaliert ihr diese etwas stärker in die Tiefe. Damit das Rad später in das Bein passt, behaltet ihr die untersten Edges selektiert und drückt `E` um sie zu extrudieren. Zieht die Edges abschließend etwas unter den Torso.






## 7. Räder
![alt text](img/7_wheels.PNG "Wheels")

Zurück in den Object-Mode `Tab`.

* Fügt für die Räder je einen Zylinder ein. Rotiert den Zylinder entlang der X-Achse, nutzt dafür den Rotate Befehl mittels der Taste `R`. Drückt also `R` gefolgt von `X` und anschließen `Num 90`. Dadruch wird der Zylinder um 90° gedreht.

![alt text](img/7_wheelsMore.PNG "More about wheels")

* Skaliert das Rad nun so, dass eines von unten in den Fuß und eines unter den Torso passt.



## 8. Duplizieren
![alt text](img/8_duplicate.PNG "Copy/Duplicate")

* Da ihr nun nur über einen Fuß und ein Rad an diesem Fuß verfügt, markiert ihr nun Rad und Fuß. Zum Auswählen mehrerer Objekte haltet ihr `Shift` gedrückt und selektiert die Objekte mit `RMB`. Drückt anschließend `Shift+D` um ein Duplikat zu erzeugen. Ein `LMB` klickt bestätigt die Aktion. Ihr könnt den Fuß und das Rad nun mit `G` an die entsprechende Stelle schieben.



## 9. Hierachie
![alt text](img/9_hierarchy.PNG "More about wheels")

* Erzeugt eine sinnvolle Hierarchie. In Blender wählt ihr immer zuerst das Child Objekt bzw. die Child Objekte und zuletzt das Objekt, dass über ihnen in der Hierarchie stehen soll (den Parent) aus. Anschließend drückt ihr `Ctrl+P` und wählt das Objekt aus.
* Euch sollte nun Rechts oben in der Szenen-Hierarchie auffallen, dass ihr eurer Parent-Objekt durch das kleine "+"-Symbol aufklappen könnt.
Die Räder sollten den Füßen, die Füße dem Torso, das Torso-Rad dem Torso und der Kopf sollte ebenfalls dem Torso untergeordnet sein.

Jetzt könnt euch der Positionsanpassungen der Beine o.ä. annehmen.

# Detailles
![alt text](img/astroBool1.PNG)

* Um eurem Droiden Detailles hinzuzufügen bietet es sich an Cubes, Spheres oder sogar komplexere Formen in die Hülle euren Astromechs zu schneiden. Hierzu verwenden wir den "Boolen Modifier".
* Als Beispiel fügen wir zunächst einen Cube mit `Shift+A` ein. Im Beispiel wurde der Cube durch Skalierung entlang der Z-Achse etwas platt gedrückt, um die Form etwas interessanter zu gestalten.

![alt text](img/astroBool2.PNG)

* Um nun die Form des Würfels in den Astromech zu schneiden wird der entsprechende Teil eures Droiden gewählt und anschließend zum Reiter mit dem Schraubschlüssel navigiert. Klickt nun unter dem Reiter auf "Add Modifier" und wählt den "Boolean Modifier".
* Wählt nun unter "Operation" die Operation "Difference" und unter Object euren "Cube" aus. Drückt anschließend "Apply". Nun könnt ihr den Cube entweder löschen oder verschieben um das Ergebnis zu begutachten.
* Der "Boolean Modifier" verfügt noch über weitere Funktionen die Formen addieren oder deren Schnittpunkt berechnen. Wundert euch jedoch nicht, wenn der Modifier nicht gut funktioniert, denn dieser ist bekannt dafür launisch zu sein (falls er klemmt: Versucht eine ähnlich große Polygonanzahl bei beiden Objekten).

![alt text](img/astroBool3.PNG)


## 11. Animation
![alt text](img/11_animation.PNG "Animation")

* Nun wird dafür gesorgt, dass dem Droiden etwas Leben eingehaucht wird.
* Wählt euren Torso aus und werft einen Blick auf die "Timeline".
* Setzt nun mit `I` gefolgt von "Location" den ersten Keyframe, an eurer Startposition.
* Wählt nun Frame 100, indem ihr an die Ensprechende Stelle klickt oder in das Feld links neben den Play-Buttons befindet. *25 Frames entsprechen standardmäßig einer Sekunde*
* zieht euren Droiden entlang der Y-Achse etwas außerhalb des Gitternetzes. Drückt nun wieder `I` und wählt "Location". Nun seht ihr auf der Timeline zwei gelbe Striche, die für die gesetzten Keyframes stehen.
* Scrollt nun etwas heraus, springt zurück zu Frame 0 und drückt "Play".
* Euer Droide fährt nun innerhalb von 4 Sekunden von Punkt A zu Punkt B.






## Aufgabe
Erstellt nun zwei weitere Astromechs mit abweichenden Formen und fügt Detailles hinzu (Antennen, Klappen, eine Linse etc.).
**Mindestanforderung:** *Unterschiedliche Beine, sichtlich geänderter Kopf, geänderte Torsoform, mind. 8 Detailles*
 
## Freiwillige Zusatzaufgabe
Verändert den Droiden so, dass sich einzelne Faces durch Extrudieren o.ä. Operationen vom Rest hervorheben und die Oberfläche interessanter gestalten.
Fügt außerdem Antennen, eine Linse und andere Detailles hinzu, damit er noch schicker aussieht.

![alt text](img/10_finished.PNG "More about wheels")


## Gelerntes

Aktion                   | Keyboard-Shortcut                   | Menübefehl 
-------------------------|-------------------------------------|------------
  Objekt-Auswahl         |   `Shift+RMB` / `B`                 | Select -> Border Select
  Box-Modelling          |                                     | 
  Loop-Selection         |   `Ctrl+Alt+RMB`                    | 
  Hierarchie (Parenting) |   `P`                               | 
  Ansichten in Blender   |   `Num 1,2,3,7` / `Ctrl+Num 1,2,3,7`| 3D-Editor -> View -> _Ansicht wählen_
  Edit vs. Object-Mode   |   `Tab`                             | 3D-Editor -> Object-Mode-Button
  Duplizieren            |  `Shift+D`                          | 3D-Editor -> Object / Mesh-> Duplicate Objects / Add Duplicate
  Grundkörper Einfügen   |  `Shift+A`                           | 3D-Editor -> Add
  Translation / Rotation / Skalierung  |  `G`(grab) / `R` / `S` | 3D-Editor -> Object / Mesh -> Transform -> _Option wählen_
  Extrudieren            |  `E`                                | 3D-Editor (Edit-Mode) -> Mesh -> Extrude -> *Option wählen*
  Edge- / Face- / Vertex-Mode  |  Edit-Mode: `Ctrl+Tab`        | 3D-Editor (Edit-Mode) -> _Edge- / Face- / Vertex-Icon_ 
  Keyframe Animation     |  `I`                                | 3D-Editor -> Toolbar (`+` links oben) -> Animation -> Keyframes -> Insert
  Alles Selektieren      |  `A`                                | 3D-Editor -> Select -> Select all by...
  Löschen                |  `X`                                | 3D-Editor -> Object / Mesh -> Delete
  Boolen Modifier        ||
