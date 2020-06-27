<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <title>Goverage</title>

  <link href="https://cdn.jsdelivr.net/npm/@mdi/font@4.x/css/materialdesignicons.min.css" rel="stylesheet">
  <link href="https://fonts.googleapis.com/css2?family=Roboto+Mono&display=swap" rel="stylesheet">

  {{range .Styles}}
  <style type="text/css">
    {{.}}
  </style>
  {{end}}

  {{range .Scripts}}
  <script>
    {{.}}
  </script>
  {{end}}



</head>

<body class="{{.Theme}}">
  <div id="app">
    <v-app id="inspire">

      <v-app-bar
        app
        dark
      >

      <v-toolbar-title>Goverage</v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn icon href="https://github.com/roger-russel/goverage">
        <v-icon>mdi-github-circle</v-icon>
      </v-btn>

      </v-app-bar>
      <v-content class="content">
        <v-container fluid >
          <v-row
              align="center"
              justify="center"
          >
            <v-col
              cols="12"
              sm="12"
              md="12"
            >

            {{ template "table-list" . }}

            </v-col>
          </v-row>


        </v-container>
      </v-content>

      <v-content class="content">
        {{ range $value := .Pages }}
          {{ template "page" $value }}
        {{ end }}
      </v-content>
    </v-app>
  </div>

  <script type="text/javascript">
    new Vue({
      el: '#app',
      vuetify: new Vuetify(),
      data() {
        return {
          search: '',
          headers: [
            {
              text: 'Files',
              align: 'start',
              sortable: false,
              value: 'name',
            },
            { text: 'Lines', value: 'lines' },
            { text: 'Green', value: 'green' },
            { text: 'Red', value: 'red' },
            { text: 'Coverage', value: 'coverage' },
          ],
          files: {{.FilesList}},
          visiblePages: {{ .VisiblePages }},
          PagesInJSON: {{ .PagesInJSON }},
        }
      },
      methods: {
        getColorClass(tracked, count) {

          if ( !tracked ) {
            return "cov-color-untracked"
          }

          if ( count < 0 ){
            count = 0
          }

          if ( count > 10 ){
            count = 10
          }

          return "cov-color-" + count

        },

        getCoverageStyle(coverage) {

          let color

          let colors = {
            0:"#df0000",
            10:"#df0d00",
            20:"#df3b00",
            30:"#df6800",
            40:"#df9200",
            50:"#ef9500",
            60:"#cddf00",
            70:"#a3df00",
            80:"#76df00",
            90:"#1bdf00",
            100:"#00df00",
          }

          if (coverage == 100) {
            color = colors[100]
          } else {
            for ( let i = 100; i => 0 ; i -= 10 ) {
              if (coverage >= i) {
                color = colors[i]
                break;
              }
            }
          }

          if (coverage == 0) {
            coverage = 100
          }

          return "background:linear-gradient(90deg, "+color+" "+coverage+"%, {{.ThemeBGColor}} "+coverage+"%);"
        },

        showPage(page){
          this.visiblePages.list[this.visiblePages.current] = false
          this.visiblePages.current = page
          this.visiblePages.list[page] = true
        },
      }
    })
  </script>

</body>
</html>
