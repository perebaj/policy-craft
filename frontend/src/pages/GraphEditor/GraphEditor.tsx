import 'reactflow/dist/style.css'

import { useCallback, useContext, useEffect, useRef, useState } from 'react'
import ReactFlow, { Background, Controls, MiniMap } from 'reactflow'

import { AddNodeEdge } from './AddNodeEdge'
import { CurrentDrawer } from './Drawers'
import { EditorProvider } from './Editor'
import { graph, GraphProvider } from './Graph'
import { generateEdge, generateNode } from './nodeGeneration'
import { allNodes } from './Nodes'
import { positionNodes } from './positionNodes'
const edgeTypes = {
  'add-node': AddNodeEdge,
}

function ReactFlowSandbox() {
  const {
    nodes,
    edges,
    reactFlowInstance,
    setReactFlowInstance,
    fitZoomToGraph,
    setNodes,
    setEdges,
  } = useContext(graph)

  const [centeredGraphAtStart, setCenteredGraphAtStart] = useState(false)
  const reactFlowRef = useRef<HTMLDivElement>(null)

  const tryCenteringGraph = useCallback(() => {
    if (centeredGraphAtStart) {
      return
    }

    fitZoomToGraph(reactFlowRef)

    const viewport = reactFlowInstance?.getViewport()
    if (viewport && viewport.x !== 0 && viewport.y !== 0) {
      return setCenteredGraphAtStart(true)
    }

    const retryTimeInMs = 50
    setTimeout(() => tryCenteringGraph(), retryTimeInMs)
  }, [centeredGraphAtStart, fitZoomToGraph, reactFlowInstance])

  useEffect(() => {
    tryCenteringGraph()
  }, [tryCenteringGraph])

  useEffect(() => {
    const initialNodes = [
      generateNode({ nodeName: 'start', id: 'start' }),
      generateNode({ nodeName: 'end' }),
    ]
    const initialEdges = [
      generateEdge({
        source: 'start',
        target: initialNodes[1].id,
      }),
    ]
    const [positionedNodes, positionedEdges] = positionNodes(
      initialNodes,
      initialEdges,
    )
    setNodes(positionedNodes)
    setEdges(positionedEdges)
  }, [setEdges, setNodes])

  return (
    <div className="relative flex h-full w-full flex-col overflow-hidden">
      <ReactFlow
        ref={reactFlowRef}
        nodes={nodes}
        edges={edges}
        edgeTypes={edgeTypes}
        nodeTypes={allNodes}
        onInit={setReactFlowInstance}
        nodesDraggable={false}
        deleteKeyCode={null}
      >
        <Background className="bg-slate-50" size={2} color="#C1C4D6" />
        <MiniMap />
        <Controls />
      </ReactFlow>
      <CurrentDrawer />
    </div>
  )
}

export function GraphEditor() {
  return (
    <EditorProvider>
      <GraphProvider>
        <ReactFlowSandbox />
      </GraphProvider>
    </EditorProvider>
  )
}
