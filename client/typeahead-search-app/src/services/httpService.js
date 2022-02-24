import axios from "axios";
import constant from '../config/constants';

function fetchSuggestions(data) {
    const findApi = constant.serverUrl + ":" + constant.serverPort + constant.findApi;
    const retData = axios
    .get(findApi, { params: { prefix: data }})
    .catch((error) => {
        console.error(error);
    });
    return retData;
}
export default fetchSuggestions;