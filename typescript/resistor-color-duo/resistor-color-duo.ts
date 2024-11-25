export const decodedValue = (arr:string[]) => {
  return COLORS.indexOf(arr[0]) * 10 + COLORS.indexOf(arr[1]);
}

export const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
]
