var fdk=require('@fnproject/fdk');

module.exports = handler;
if(process.env['FN_PATH'] !== undefined) {
  fdk.handle(handler);
}

function handler(input){
  var name = 'World';
  if (input.name) {
    name = input.name;
  }
  response = {'message': 'Hello ' + name}
  return response
};
