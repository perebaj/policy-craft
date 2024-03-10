import { Handle, NodeProps, Position } from 'reactflow'

import { NodeWrapper } from './NodeWrapper'

type DecisionNodeData = {
  width: number
  height: number
}

// DecisionNode can't be created by the user, only automatically. It represents the failure of a conditional node.
export function DecisionNode({ data }: NodeProps<DecisionNodeData>) {
  return (
    <NodeWrapper>
      <div
        className={`flex h-full items-center justify-center rounded-lg border-4 border-red-300 bg-white text-[12px]`}
        style={{
          width: data.width * 1.5,
          height: data.height,
        }}
      >
        <Handle
          type="target"
          id="decision"
          className="invisible"
          position={Position.Top}
          isConnectable={false}
        />
        <p className="font-medium">{'Decision'}</p>
      </div>
    </NodeWrapper>
  )
}
