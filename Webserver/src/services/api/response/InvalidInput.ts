export function InvalidInput(res, message:String) {
    return res.status(200).json({
        status: 200,
        message: message
    });
}