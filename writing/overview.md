# Overview of "Untitled Summer"

## Notes in Progress for documentation and explanations later

DBL 610 -- Kate Folkenroth (Global Health Studies & Integrative Informatics)

### 6.2.2022

- Working on standards/decisions
    - [File type](#file-type)
        - TEI-XML or RDF
    - [Metadata standards](#metadata-standards)
        - DublinCore
    - [Coding languages](#coding-languages)
        - GoLang (why)
            binary library, netlify 
        - Html/CSS for static sites
    - Where?
        - GitHub 
            - Make an Organization?
        - Allegheny Server
        - Netlify? Hugo?
        - (etc)
    
    #### File Type 

    * RDF/XML seems likely
    > xmls 

    ```xml vs rdf
    XML is a syntax while RDF is a data model

    RDF documents are written in XML. The XML language used by RDF is called RDF/XML. By using XML, RDF information can easily be exchanged between different types of computers using different types of operating systems and application languages.
    ```

    ```tei xml vs rdf xml
        TEI is an XML-based markup language that enable scholars to store, analyze, and share humanities textual information. The Text Encoding Initiative (TEI) Consortium is a organization that develops the standards for text encoding in the TEI encoding language.

        However, TEI is XML based, and thus it suffers heavy limitations, such as overlapping annotations. RDF is able to overcome this, using ontology vocabularies and allowing powerful queries.

        http://rdftef.sourceforge.net/
        RdfTef >> RDF Textual Encoding Framework
        is an open source Java framework that supports textual encoding in RDF+OWL. It provides a shell that allows the model to be queried using SPARQL, is able to import existing encoded text in XML TEI format, and can save the model to RDF/XML or export it in form of aspect slices, as XML TEI format.

    ```

    ```Sparql
    SPARQL (pronounced "sparkle" /ˈspɑːkəl/, a recursive acronym for SPARQL Protocol and RDF Query Language) is an RDF query language
    ```

    ```Otro Links
    http://rdftef.sourceforge.net/
    >> https://github.com/knakk/sparql

    > consider database layer on archive to make static pages eventually

    headless cms vs static site gen
    ```


    // !ch3 notes
    ```
    !dead resource > not all research can be public 
    who is the custodian of records // permission-ing
    > someone asks for access > determined by custodian > given identifier/key
        will they publish them? ethics?

    this project is created only in respect to publicly released materials
    technical solution
        point is to release something publicly > private resources are stored elsewhere
    ```

    === 
    ``` general xml rdf resources
    https://www.w3schools.com/xml/xml_services.asp  >> general
    https://www.w3schools.com/xml/default.asp >> study xml

    https://www.dublincore.org/specifications/dublin-core/usageguide/#rdfxml
        https://www.dublincore.org/specifications/dublin-core/dc-xml-guidelines/
        https://www.dublincore.org/specifications/dublin-core/dcmes-xml/

    ```

    #### Metadata Standards
    ```Dublin bb
    The Dublin Core™ is a set of fifteen basic categories (such as creator, title, subject, and publisher) for describing information resources (see http://dublincore.org/)

    https://www.dublincore.org/specifications/dublin-core/dcmes-xml/

    ```

        ***
    make a maximalist record >> flesh out a template
    > then something that can parse it 
    (dot notation? item.title)
    provenance ---  item.ownership.3

    * [Creating Metadata](https://www.dublincore.org/resources/userguide/creating_metadata/)

    #### Coding Languages

    /// recursion to parse through nodes of arrays
        arrays & nodes >
    make xml parsable

    //end of the line boi


    ##### Fun with words
    - folk.chompe.rs
        - kate.folk >> Kate Folkenroth
    - folkArchive
    - Allegheny Archives // Allegheny Archivists
        - a collection of items donated and/or displayed on the campus
        - purpose: to use for education
    - Initiative: Curating the College
    - An Institutional Repository 
    - 


//
use mustache


template
 the data will be throw


 html w/ mustaches populated

 ===

 reading the implementation 

 code in the parsing file to call things and then make things



partials, for header and footers > folder structure 
tempalate and partials and then load em all

>> make themes 

html w/ mustache > item w/ 

musatche loops 
non-exhaustive variables

```
<ul>
{{{names}}}
<li>{{{.}}}</li>
{{{/names}}}
</ul>
```

LUMAN-- 
I cant run mustache on my machine im just having a time 
