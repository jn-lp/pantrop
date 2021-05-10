import isInDOM from './isInDom'

export default function hasParent(element: any, root: any) {
  return root && root.contains(element) && isInDOM(element)
}
