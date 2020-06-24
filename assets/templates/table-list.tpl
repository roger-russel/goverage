{{ define "table-list" }}
<v-card class="table-list">
  <v-card-title>
    <v-spacer></v-spacer>
    <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details>
    </v-text-field>
  </v-card-title>
  <v-data-table
    :headers="headers"
    :items="files"
    :search="search"
  >

  <template v-slot:body="{ items }">
    <tbody>
      <tr v-for="item in items" :key="item.name">
        <td>
          {{"{{"}} item.name {{"}}"}}
        </td>
        <td>
          {{"{{"}} item.lines {{"}}"}}
        </td>
        <td>
          {{"{{"}} item.green {{"}}"}}
        </td>
        <td>
          {{"{{"}} item.red {{"}}"}}
        </td>
        <td
          class="coverage"
          :style="getCoverageStyle(item.coverage)"
        >
          {{"{{"}} item.coverage.toFixed(2) {{"}}"}} %
        </td>
      </tr>
    </tbody>
  </template>
  </v-data-table>
</v-card>
{{ end }}
