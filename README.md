# GoWebStuff
Aktuelle Entwicklungen im Bereich Online Medien - Evaluierung von Go für Webdevelopment

Ordnerstruktur:

- **Executable**
- courseconfig.json
- info.md
- 🗀 templates
- 🗀 static
    - 🗀 css
    - 🗀 js
    - 🗀 highlightJS
    - 🗀 bootstrap
- 🗀 Userdata
    - 🗀 Portraits
        * ...[Matrikelnumer].png
        * default.png
    - 🗀 Students
        - ...🗀 [Matrikelnummer]
            - profile.json
            - post_[postnumber].md
            - post_001.md
            - post_002.md
            - post_...

**[wkhtmltopdf](https://wkhtmltopdf.org/) muss für die PDF Funktionalität installiert und zu PATH hinzugefügt sein.**

- Max GroupNumber: 6

Da Bilder in den Aufgabenstellungen von der Haupseite aus aufgerufen werden, können diese per Markdown folgendermaßen z.b. aufgerufen werden: ![beschreibung](img/06.png "Licht"). Da der assignments/img Ordner beliebig uterteilt werden kann, ist z.b. auch ![alt text](img/a1/06.png "Licht") möglich - wenn das Bild in einem Unterordner assignments/img/a1 abgelegt wurde.

