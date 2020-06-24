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
        color="cyan"
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
              <v-card>
                <v-card-title>
                  <v-spacer></v-spacer>
                  <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details>
                  </v-text-field>
                </v-card-title>
                <v-data-table :headers="headers" :items="files" :search="search"></v-data-table>
              </v-card>
            </v-col>
          </v-row>


        </v-container>
      </v-content>

      <v-content class="content">
        {{ range .Pages}}
          {{ template "page" . }}
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

        }
      }
    })
  </script>

</body>
</html>
