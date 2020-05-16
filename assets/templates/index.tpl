<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <title>Goverage</title>

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

<body>
  <div id="app">
    <v-app id="inspire">
      <v-card>
        <v-card-title>
          Goverage
          <v-spacer></v-spacer>
          <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details>
          </v-text-field>
        </v-card-title>
        <v-data-table :headers="headers" :items="files" :search="search"></v-data-table>
      </v-card>
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
          files: [
            {
              name: 'Frozen Yogurt',
              lines: 159,
              green: 6.0,
              red: 4.0,
              coverage: '1%',
            },
          ],
        }
      },
    })
  </script>

</body>
</html>
