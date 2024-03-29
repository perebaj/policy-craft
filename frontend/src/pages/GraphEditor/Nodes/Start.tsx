import { Handle, NodeProps, Position } from 'reactflow'

import { NodeWrapper } from './NodeWrapper'

type StartNodeData = {
  width: number
  height: number
}

export function StartNode({ data }: NodeProps<StartNodeData>) {
  return (
    <NodeWrapper>
      <div
        className={`flex aspect-square h-full items-center justify-center rounded-full border-4 border-slate-300 bg-white text-[12px]`}
        style={{
          width: data.width,
          height: data.height,
        }}
      >
        <p className="font-medium">{'Start'}</p>
        <Handle
          type="source"
          id="source"
          className="invisible"
          position={Position.Bottom}
          isConnectable={false}
        />
      </div>
    </NodeWrapper>
  )
}
