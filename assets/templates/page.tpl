{{ define "page" }}
<v-container fluid class="page" v-if="visiblePages.list['{{ .FullName }}']" >
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
        Page: {{ .FullName }}
        </v-card-title>
          <v-simple-table
            dense
          >
            <template v-slot:default>
              <tbody>
              {{ range .Lines }}
                <tr>
                  <td
                    title="{{ .Line }}"
                    class="line"
                    width="50px"
                  >{{ .Line }}</td>
                  <td>
                    {{ range .Contents }}
                      <span
                      :class="['pre', getColorClass({{.Tracked}},{{.Count}})]"
                      title="{{.Count}}"
                      >{{ .Content }}</span>
                    {{ end}}
                  </td>
                </tr>
              {{ end }}
              </tbody>
            </template>
          </v-simple-table>
      </v-card>
    </v-col>
  </v-row>
</v-container>
{{ end }}
