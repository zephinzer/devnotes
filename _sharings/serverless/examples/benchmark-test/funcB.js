var fdk=require('@fnproject/fdk');

module.exports = handler;
if(process.env['FN_PATH'] !== undefined) {
  fdk.handle(handler);
}

function handler(input){
  return `Hello ${input ? input.name ? input.name : 'World' : 'World'}`;
};
