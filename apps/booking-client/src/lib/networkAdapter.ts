export const enum NetworkAdapter {
    POST = 'POST',
    GET = 'GET',
    PUT = 'PUT',
    DELETE = 'DELETE',
}

export const NewDefaultHeader = () => {
    const newHeaders = new Headers();
    return newHeaders;
}