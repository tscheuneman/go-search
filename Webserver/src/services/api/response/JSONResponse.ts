export function JSONResponse(res, payload:Object) {
    return res.status(200).json({
        status: 200,
        payload: payload
    });
}