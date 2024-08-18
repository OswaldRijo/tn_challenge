import { initReactI18next } from 'react-i18next';
import i18next from 'i18next';

import engtable from './en/client.json';
import engServer from './en/server.json';
import estable from './es/client.json';
import esServer from './es/server.json';

i18next.use(initReactI18next).init({
  lng: 'en', // if you're using a language detector, do not define the lng option
  debug: true,
  resources: {
    en: {
      client: engtable,
      server: engServer
    },
    es: {
      client: estable,
      server: esServer
    }
  },
  ns: ['client', 'server'],
  defaultNS: 'form'
});
