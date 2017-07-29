/**
 *
 * This file is used to provide type hinting in your javascript plugins.
 * It is as well a documentation which functions/global vars you can use.
 * It is not used in any other way and is never called.
 *
 */


/**
 * The name of the user who triggered the event
 */
var sender;

/**
 * The name of the channel the event was triggered in
 */
var channel;

/**
 * Everything the user wrote after the initial command
 */
var params;

/**
 * Sends a message to the channel where the current event was triggered in
 * @param {string} message
 */
function sendMessage(message) {
}

/**
 * Returns a user data object as json string.
 * Use JSON.parse() to get a json object.
 * @param {string} name
 * @returns {string}
 */
function getUser(name) {
    return {
        Name: '',
        MessageCount: 0,
        LastJoin: '',
        LastPart: '',
        LastActive: '',
        FirstSeen: ''
        };
}

/**
 * Save json data by key.
 * All data is saved in a plugin namespace
 * @param {string} key
 * @param {string} data
 */
function setData(key, data) {

}

/**
 * Get data from your plugin namespace
 * @param {string} key
 * @returns {string}
 */
function getData(key) {
    return "{}"
}