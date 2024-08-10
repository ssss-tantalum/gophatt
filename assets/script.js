import 'htmx.org'
import Alpine from 'alpinejs'
import { persist } from '@alpinejs/persist'

window.Alpine = Alpine
Alpine.plugin(persist)

Alpine.store('darkMode', {
  on: Alpine.$persist(true).as('darkMode'),
  toggle: function () {
    this.on = !this.on
  },
})

Alpine.start()
