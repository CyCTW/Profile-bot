import axios from "axios";
 
const services = {};

const instance = axios.create({
  baseURL: "",
  withCredentials: true
});

export const createActivity = ({activity, date, place, idToken}) => {
    return instance.post(`/activity/${idToken}`, {
        activity,
        date,
        place,
    })
}
