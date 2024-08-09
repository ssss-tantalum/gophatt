import 'htmx.org'
import Alpine from 'alpinejs'

window.Alpine = Alpine

Alpine.store('darkMode', {
  on: false,
  toggle() {
    this.on = !this.on
  },
})

Alpine.start()
