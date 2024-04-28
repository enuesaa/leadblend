type Note = {
  name: string
  content: string
}
export const fetchNotes = async (): Promise<Note[]> => {
  const list = [
    {
      name: 'タイトル',
      content: '内容',
    },
  ]
  return list
}
