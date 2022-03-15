import EventBus from './eventbus';
import { EVENTS } from '../constants';

type ApiMethod = "GET" | "POST";
interface ApiOptions {
    method?: ApiMethod,
    body?: any;
    headers?: Record<string, string>;
}

const DEFAULT_SETTINGS = {
    method: 'GET'
}

export const ApiRequest = (path: string, callback: (response: any) => void, options?: ApiOptions,) => {
    fetch(path, { ...DEFAULT_SETTINGS, ...options }).then(res => {
        if(res.status === 200) {
            res.json().then(response => {
                console.log(response);
                callback(response);
            });
        } else {
            switch(res.status) {
                case 403:
                    EventBus.trigger(EVENTS.NAVIGATE, '/login');
                    break;
            }
        }

    }).catch(err => {
        EventBus.trigger(EVENTS.API_ERROR)
        console.error(err)
    });
}
