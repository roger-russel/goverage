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
            { text: 'Yellow', value: 'yellow' },
            { text: 'Red', value: 'red' },
            { text: 'Coverage', value: 'coverage' },
          ],
          files: [
            {
              name: 'Frozen Yogurt',
              lines: 159,
              green: 6.0,
              yellow: 24,
              red: 4.0,
              coverage: '1%',
            },
            {
              name: 'Ice cream sandwich',
              lines: 237,
              green: 9.0,
              yellow: 37,
              red: 4.3,
              coverage: '1%',
            },
            {
              name: 'Eclair',
              lines: 262,
              green: 16.0,
              yellow: 23,
              red: 6.0,
              coverage: '7%',
            },
            {
              name: 'Cupcake',
              lines: 305,
              green: 3.7,
              yellow: 67,
              red: 4.3,
              coverage: '8%',
            },
            {
              name: 'Gingerbread',
              lines: 356,
              green: 16.0,
              yellow: 49,
              red: 3.9,
              coverage: '16%',
            },
            {
              name: 'Jelly bean',
              lines: 375,
              green: 0.0,
              yellow: 94,
              red: 0.0,
              coverage: '0%',
            },
            {
              name: 'Lollipop',
              lines: 392,
              green: 0.2,
              yellow: 98,
              red: 0,
              coverage: '2%',
            },
            {
              name: 'Honeycomb',
              lines: 408,
              green: 3.2,
              yellow: 87,
              red: 6.5,
              coverage: '45%',
            },
            {
              name: 'Donut',
              lines: 452,
              green: 25.0,
              yellow: 51,
              red: 4.9,
              coverage: '22%',
            },
            {
              name: 'KitKat',
              lines: 518,
              green: 26.0,
              yellow: 65,
              red: 7,
              coverage: '6%',
            },
          ],
        }
      },
    })
  </script>

</body>
</html>
