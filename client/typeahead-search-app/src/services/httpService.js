import axios from "axios";

function fetchSuggestions(data) {
    const retData = axios
    .get("http://localhost:4000/find", { params: { prefix: data }})
    .catch((error) => {
        console.error(error);
    });
    return retData;
}
export default fetchSuggestions;