
# Evago - Evaluierung von Go für Webdevelopment

### Veranstaltung: Aktuelle Entwicklungen im Bereich Online Medien - Prof. Eisenbiegler

### Beispielhafte Ordnerstruktur:

- 🗀 Hauptverzeichniss
    - **Executable** (Evalgo.exe / Evalgo linux binary)
    - courseconfig.json
    - info.md
    - 🗀 templates
    - 🗀 static
        - 🗀 css
        - 🗀 js
        - 🗀 highlightJS
        - 🗀 bootstrap
    - 🗀 coursedata
        - 🗀 tasks
            - 🗀 img (beliebige Unterstruktur)
                - 🗀 unterordner1
                    - bild.png
                - coolesBild01.png
            - post_001.md
            - post_002.md
            - post_...
        - 🗀 portraits
            * ...[Matrikelnumer].png
            * default.png
        - 🗀 users
            - ...🗀 [Matrikelnummer]
                - profile.json
                - post_[postnumber].md
                - post_001.md
                - post_002.md
                - post_...

### Konfiguration
Im Hauptverzeichniss der Anwendung muss eine Datei `courseconfig.json` liegen. Der Inhalt sieht folgendermaßen aus:
```json
{
    "port": ":8080",
    "course_name": "Evalgo",
    "group_number": 4,
    "open_course": false,
    "root_url": "/evalgo/",
    "master_password": "gulasch",
    "tutors_can_post": true,
    "enable_grades": true,
    "enable_cards": true,
    "classes":[ "MIB", "MKB", "OMB" ]
}
```
- `port` bestimmt den Port auf dem Server, auf dem die Anwendung laufen soll.
- `course_name` bestimmt den Titel des Kurses, der auf der Hauptseite angezeigt wird.
- `group_number` bestimmt die Anzahl der Gruppen (Farben), in die die Kursteilnehmer unterteilt werden.
- `open_course` bestimmt, ob auch Studenten die Beiträge ihrer Mitstudierenden sehen können, oder nur Tutoren und Admins.
- `root_url` bestimmt die URL der Haupseite. Hierbei ist zu beachten, dass die Schrägstriche vorne und hinten benötigt werden.
- `tutors_can_post` Ermöglicht es Tutoren auch Abgaben zu machen oder nicht.
- `enable_grades` Schaltet beim Feedback die Notenvergabe frei
- `enable_cards` Schaltet beim Feedback die vergabe von gelben / roten Karten frei.

Die Variablen `master_password`, und `classes` sind noch ungenutzt und dienen als Platzhalter für zukünftige Versionen. Die Anwendung muss neu gestartet werden, damit Änderungen an der Konfigurationsdatei berücksichtigt werden.

### Nutzergruppen
Es gibt drei Nutzergruppen:

- 1 `Student`
- 2 `Tutor`
- 3 `Admin`

Momentan werden alle Nutzer als `Student` registriert. In der entsprechenden profile.json Datei kann im nachhinein dann "usertype" auf 2 (`Tutor`) oder 3 (`Admin`) gesetzt werden. Momentan gibt es in der Funktionalität noch keinen unterschied zwischen Tutor und Admin.

### Aufgabenstellung
Für die Aufgabenstellung werden Markdown Dateien im coursedata/tasks erstellt. Diese müssen dem Benennungsschema `post_001.md`, `post_002.md`... folgen. Der Aufgabenwähler auf der Hauptseite erkennt automatisch, für welche Aufgaben Einträge bestehen und zeigt diese im Aufgabenwähler an.

Da Bilder in den Aufgabenstellungen von der Haupseite aus aufgerufen werden, können diese per Markdown folgendermaßen z.b. aufgerufen werden: `![beschreibung](img/06.png "Licht")`. 
Da der assignments/img Ordner beliebig unterteilt werden kann, ist z.b. auch `![alt text](img/a1/06.png "Licht")` möglich - wenn das Bild in einem Unterordner assignments/img/a1 abgelegt wurde.

### Posts
Über den Reiter "Post" können Studierende Abgaben machen.

### Feedback
Für jede ausgewählte Abgabe können Tutoren und Admins Feedback in Textform - und je nach konfiguration auch mit Noten und roten / gelben Karten geben.

### PDF Download
**[wkhtmltopdf](https://wkhtmltopdf.org/) muss für die PDF Funktionalität installiert und zu PATH hinzugefügt sein.**
Dann kann über die URL `.../[Matrikelnummer]/pdf` eine PDF Datei mit allen Abgaben des in der URL mit der Matrikelnummer bestimmten Studenten erstellen. Download dieser Dateien ist momentan noch nicht möglich, aber geplant.
