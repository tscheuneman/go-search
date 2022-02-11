import axios from 'axios';
export const apiCall = (options:object, url:string, data:object, method: string, postParams: boolean = false, pdfDownload = false) => {
    return new Promise((resolve, reject) => {
        let requestOptions:object = {url, ...options, method};
        if(!postParams) {
            switch(method) {
                case 'GET': case 'get':
                    (requestOptions as any).params = data;
                    break;
                default:
                    (requestOptions as any) = data;
            }
        } else {
            switch(method) {
                case 'GET': case 'get':
                    (requestOptions as any).params = data;
                    break;
                default:
                    (requestOptions as any).params = data;
            }
        }

        if(pdfDownload) {
            axios.request(requestOptions).then(response => {
                if(!response) {
                    reject('An unknown error has occured');
                } else {
                    resolve(response.data);
                }
            }).catch(error => {
                reject(error);
            });
        } else {
            axios.request(requestOptions).then(response => {
                if(response && typeof(response.data) !== 'object' && typeof(response.data) !== 'string') {
                    reject('Invalid Response');
                } else if(!response) {
                    reject('An unknown error has occured');
                } else if(response) {
                    resolve(response.data);
                }
            }).catch(error => {
                reject(error);
            });
        }

    });
};
